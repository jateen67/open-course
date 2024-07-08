import datetime
from pydantic import BaseModel


class NotificationBase(BaseModel):
    time_sent: datetime


class NotificationCreate(NotificationBase):
    pass


class Notification(NotificationBase):
    id: int
    FK_Order: int
    FK_Notification_Type: int

    class Config:
        orm_mode = True


class NotificationTypeBase(BaseModel):
    type: str


class NotificationTypeCreate(NotificationTypeBase):
    pass


class NotificationType(NotificationTypeBase):
    id: int
    notifications: list[Notification] = []

    class Config:
        orm_mode = True


class OrderBase(BaseModel):
    is_active: bool
    updated_at: datetime


class OrderCreate(OrderBase):
    pass


class Order(OrderBase):
    id: int
    FK_User: int
    FK_Course: int
    notifications: list[Notification] = []

    class Config:
        orm_mode = True


class CourseBase(BaseModel):
    course_code: str
    course_title: str
    open_seats: int
    waitlist_available: int
    waitlist_capacity: int
    created_at: datetime
    updated_at: datetime


class CourseCreate(CourseBase):
    pass


class Course(CourseBase):
    id: int
    orders: list[Order] = []

    class Config:
        orm_mode = True


class UserBase(BaseModel):
    name: str
    email: str
    phone: int
    created_at: datetime


class UserCreate(UserBase):
    pass


class User(UserBase):
    id: int
    orders: list[Order] = []

    class Config:
        orm_mode = True
