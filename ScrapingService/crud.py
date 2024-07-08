from sqlalchemy.orm import Session

from . import models, schemas


def get_users(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.User).offset(skip).limit(limit).all()


def get_user(db: Session, user_id: int):
    return db.query(models.User).filter(models.User.id == user_id).first()


def get_user_by_email(db: Session, email: str):
    return db.query(models.User).filter(models.User.email == email).first()


def get_user_by_phone(db: Session, phone: int):
    return db.query(models.User).filter(models.User.phone == phone).first()


def create_user(db: Session, user: schemas.UserCreate):
    db_user = models.User(
        name=user.name,
        email=user.email,
        phone=user.phone,
    )
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user


def get_courses(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.Course).offset(skip).limit(limit).all()


def get_course(db: Session, course_id: int):
    return db.query(models.Course).filter(models.Course.id == course_id).first()


def get_course_by_course_code(db: Session, course_code: str):
    return (
        db.query(models.Course).filter(models.Course.course_code == course_code).first()
    )


def get_course_by_course_title(db: Session, course_title: str):
    return (
        db.query(models.Course)
        .filter(models.Course.course_title == course_title)
        .first()
    )


def create_course(db: Session, course: schemas.CourseCreate):
    db_course = models.Course(
        course_code=course.course_code,
        course_title=course.course_title,
        open_seats=course.open_seats,
        waitlist_available=course.waitlist_available,
        waitlist_capacity=course.waitlist_capacity,
    )
    db.add(db_course)
    db.commit()
    db.refresh(db_course)
    return db_course


def update_course(db: Session):
    raise NotImplementedError("This method must be implemented")


def get_orders(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.Order).offset(skip).limit(limit).all()


def get_order(db: Session, order_id: int):
    return db.query(models.Order).filter(models.Order.id == order_id).first()


def get_order_by_user_id(db: Session, user_id: str):
    return db.query(models.Order).filter(models.Order.FK_User == user_id).first()


def get_order_by_course_id(db: Session, course_id: str):
    return db.query(models.Order).filter(models.Order.FK_Course == course_id).first()


def create_order(db: Session, order: schemas.OrderCreate, user_id: int, course_id: int):
    db_order = models.Order(
        **order.dict(),
        FK_User=user_id,
        FK_Course=course_id,
    )
    db.add(db_order)
    db.commit()
    db.refresh(db_order)
    return db_order


def update_order(db: Session):
    raise NotImplementedError("This method must be implemented")


def get_notifications(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.Notification).offset(skip).limit(limit).all()


def get_notification(db: Session, notification_id: int):
    return (
        db.query(models.Notification)
        .filter(models.Notification.id == notification_id)
        .first()
    )


def get_notification_by_order_id(db: Session, order_id: str):
    return (
        db.query(models.Notification)
        .filter(models.Notification.FK_Order == order_id)
        .first()
    )


def create_notification(
    db: Session,
    notification: schemas.NotificationCreate,
    order_id: int,
    notification_type_id: int,
):
    db_notification = models.Notification(
        **notification.dict(),
        FK_Order=order_id,
        FK_Notification_Type=notification_type_id,
    )
    db.add(db_notification)
    db.commit()
    db.refresh(db_notification)
    return db_notification


def get_notification_types(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.Notification_Type).offset(skip).limit(limit).all()
