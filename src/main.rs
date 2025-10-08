#[macro_use] extern crate rocket;

#[launch]
async fn rocket() -> _ {
    let database_url = std::env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    println!("Connecting to database at: {}", database_url);
    let db_pool = sqlx::SqlitePool::connect(&database_url).await.expect("Failed to connect to database");

    rocket::build()
        .manage(db_pool)
        .mount("/", routes![quote])
}

#[derive(sqlx::FromRow, serde::Serialize)]
#[derive(Debug)]
struct Quote {
    id: i64,
    quote: String,
    author: String,
}

#[derive(rocket::form::FromFormField)]
enum UnitOfTime {
    Second,
    Minute,
    Hour,
    Day,
    Week,
    Fortnight,
    Month,
    Year,
}

use rocket::serde::json::Json;
use rocket::http::Status;
use chrono::{Datelike, Timelike, Utc, NaiveDate, NaiveDateTime, Duration};
    
#[get("/quote?<unit_of_time>")]
async fn quote(unit_of_time: Option<UnitOfTime>, conn: &rocket::State<sqlx::SqlitePool>) -> Result<Json<Quote>, Status> {
    match query_random_quote(conn)
        .await
    {   
        Ok(quote) => Ok(Json(quote)),
        Err(e) => {
            eprintln!("Database error: {}", e);
            Err(Status::InternalServerError)
        }
    }
}

async fn query_random_quote(conn: &sqlx::SqlitePool) -> Result<Quote, sqlx::Error> {
    sqlx::query_as::<_, Quote>("SELECT id, quote, author FROM quotes ORDER BY RANDOM() LIMIT 1")
        .fetch_one(conn)
        .await
}

async fn query_quote_count(conn: &sqlx::SqlitePool) -> Result<i64, sqlx::Error> {
    let (count,): (i64,) = sqlx::query_as("SELECT COUNT(*) FROM quotes")
        .fetch_one(conn)
        .await?;
    Ok(count)
}

async fn query_quote_by_id(conn: &sqlx::SqlitePool, id: i64) -> Result<Quote, sqlx::Error> {
    sqlx::query_as::<_, Quote>("SELECT id, quote, author FROM quotes WHERE id = ?")
        .bind(id)
        .fetch_one(conn)
        .await
}

fn unit_of_time_to_epoch(unit_of_time: UnitOfTime) -> Result<i64, &'static str> {
    let now = Utc::now();
    let date = now.date_naive();

    let date_time = match unit_of_time {
        UnitOfTime::Second => NaiveDateTime::new(
            date,
            chrono::NaiveTime::from_hms_opt(now.hour(), now.minute(), now.second())
                .ok_or("Invalid time for Second")?,
        ),
        UnitOfTime::Minute => NaiveDateTime::new(
            date,
            chrono::NaiveTime::from_hms_opt(now.hour(), now.minute(), 0)
                .ok_or("Invalid time for Minute")?,
        ),
        UnitOfTime::Hour => NaiveDateTime::new(
            date,
            chrono::NaiveTime::from_hms_opt(now.hour(), 0, 0)
                .ok_or("Invalid time for Hour")?,
        ),
        UnitOfTime::Day => NaiveDateTime::new(
            date,
            chrono::NaiveTime::from_hms_opt(0, 0, 0)
                .ok_or("Invalid time for Day")?,
        ),
        UnitOfTime::Week => {
            let weekday = date.weekday().num_days_from_monday();
            let start_of_week = date - Duration::days(weekday.into());
            NaiveDateTime::new(
                start_of_week,
                chrono::NaiveTime::from_hms_opt(0, 0, 0)
                    .ok_or("Invalid time for Week")?,
            )
        }
        UnitOfTime::Fortnight => {
            let days_since_fortnight = (date.ordinal() - 1) % 14;
            let start_of_fortnight = date - Duration::days(days_since_fortnight.into());
            NaiveDateTime::new(
                start_of_fortnight,
                chrono::NaiveTime::from_hms_opt(0, 0, 0)
                    .ok_or("Invalid time for Fortnight")?,
            )
        }
        UnitOfTime::Month => NaiveDateTime::new(
            NaiveDate::from_ymd_opt(date.year(), date.month(), 1)
                .ok_or("Invalid date for Month")?,
            chrono::NaiveTime::from_hms_opt(0, 0, 0)
                .ok_or("Invalid time for Month")?,
        ),
        UnitOfTime::Year => NaiveDateTime::new(
            NaiveDate::from_ymd_opt(date.year(), 1, 1)
                .ok_or("Invalid date for Year")?,
            chrono::NaiveTime::from_hms_opt(0, 0, 0)
                .ok_or("Invalid time for Year")?,
        ),
    };

    Ok(date_time.and_utc().timestamp())
}
