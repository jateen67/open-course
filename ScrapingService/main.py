from fastapi import Depends, FastAPI, HTTPException
from sqlalchemy.orm import Session

from . import crud, models, schemas
from .database import SessionLocal, engine

models.Base.metadata.create_all(bind=engine)

app = FastAPI()


# Dependency
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


@app.get("/users/", response_model=list[schemas.User])
def get_users(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    return crud.get_users(db, skip=skip, limit=limit)


@app.get("/users/{user_id}", response_model=schemas.User)
def get_user(user_id: int, db: Session = Depends(get_db)):
    db_user = crud.get_user(db, user_id=user_id)
    if db_user is None:
        raise HTTPException(status_code=404, detail="User not found")
    return db_user


@app.post("/users/", response_model=schemas.User)
def create_user(user: schemas.UserCreate, db: Session = Depends(get_db)):
    db_user = crud.get_user_by_email(db, email=user.email)
    if db_user:
        raise HTTPException(status_code=400, detail="Email already registered")
    return crud.create_user(db=db, user=user)


@app.get("/courses/", response_model=list[schemas.Course])
def get_courses(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    return crud.get_courses(db, skip=skip, limit=limit)


@app.get("/courses/{course_id}", response_model=schemas.Course)
def get_course(course_id: int, db: Session = Depends(get_db)):
    db_user = crud.get_course(db, course_id=course_id)
    if db_user is None:
        raise HTTPException(status_code=404, detail="Course not found")
    return db_user


@app.post("/courses/", response_model=schemas.Course)
def create_course(course: schemas.CourseCreate, db: Session = Depends(get_db)):
    db_course = crud.get_course_by_course_code(db, course=course.course_code)
    if db_course:
        raise HTTPException(status_code=400, detail="Course already in database")
    return crud.create_course(db=db, course=course)


@app.put("/courses/", response_model=schemas.Course)
def update_course():
    raise HTTPException(status_code=400, detail="To implement")


@app.get("/orders/", response_model=list[schemas.Order])
def get_orders(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    return crud.get_orders(db, skip=skip, limit=limit)


@app.get("/orders/{order_id}", response_model=schemas.Order)
def get_order(order_id: int, db: Session = Depends(get_db)):
    db_order = crud.get_order(db, order_id=order_id)
    if db_order is None:
        raise HTTPException(status_code=404, detail="Order not found")
    return db_order


@app.post("/orders/", response_model=schemas.Order)
def create_order(
    order: schemas.OrderCreate,
    user_id: int,
    course_id: int,
    db: Session = Depends(get_db),
):
    return crud.create_order(db=db, order=order, user_id=user_id, course_id=course_id)


@app.put("/orders/", response_model=schemas.Order)
def update_order():
    raise HTTPException(status_code=400, detail="To implement")


@app.get("/notifications/", response_model=list[schemas.Notification])
def get_notifications(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    return crud.get_notifications(db, skip=skip, limit=limit)


@app.get("/notifications/{notification_id}", response_model=schemas.Notification)
def get_notification(notification_id: int, db: Session = Depends(get_db)):
    db_notification = crud.get_notification(db, notification_id=notification_id)
    if db_notification is None:
        raise HTTPException(status_code=404, detail="Notification not found")
    return db_notification


@app.post("/notifications/", response_model=schemas.Notification)
def create_notification(
    notification: schemas.NotificationCreate,
    order_id: int,
    notification_type_id: int,
    db: Session = Depends(get_db),
):
    return crud.create_notification(
        db=db,
        notification=notification,
        order_id=order_id,
        notification_type_id=notification_type_id,
    )
