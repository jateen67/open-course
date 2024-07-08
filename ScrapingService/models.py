import datetime
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from sqlalchemy.orm import relationship
from .database import Base


class User(Base):
    __tablename__ = "tbl_Users"

    id = Column(Integer, primary_key=True)
    name = Column(String)
    email = Column(String)
    phone = Column(Integer, index=True)
    created_at = Column(DateTime, default=datetime.datetime.now())

    orders = relationship("Order", back_populates="user")


class Course(Base):
    __tablename__ = "tbl_Courses"

    id = Column(Integer, primary_key=True)
    course_code = Column(String, unique=True, index=True)
    course_title = Column(String)
    open_seats = Column(Integer)
    waitlist_available = Column(Integer)
    waitlist_capacity = Column(Integer)
    created_at = Column(DateTime, default=datetime.datetime.now())
    updated_at = Column(DateTime, default=datetime.datetime.now())

    orders = relationship("Order", back_populates="course")


class Order(Base):
    __tablename__ = "tbl_Orders"

    id = Column(Integer, primary_key=True)
    FK_User = Column(Integer, ForeignKey("tbl_Users.id"))
    FK_Course = Column(Integer, ForeignKey("tbl_Courses.id"))
    is_active = Column(Boolean, default=True, index=True)
    updated_at = Column(DateTime, default=datetime.datetime.now())

    user = relationship("User", back_populates="orders")
    course = relationship("Course", back_populates="orders")
    notifications = relationship("Notification", back_populates="order")


class Notification(Base):
    __tablename__ = "tbl_Notifications"

    id = Column(Integer, primary_key=True)
    FK_Order = Column(Integer, ForeignKey("tbl_Orders.id"))
    FK_Notification_Type = Column(Integer, ForeignKey("tbl_Notification_Types.id"))
    time_sent = Column(DateTime, default=datetime.datetime.now())

    order = relationship("Order", back_populates="notifications")
    notification_type = relationship(
        "Notification_Type", back_populates="notifications"
    )


class Notification_Type(Base):
    __tablename__ = "tbl_Notification_Types"

    id = Column(Integer, primary_key=True)
    type = Column(String)

    notifications = relationship("Notification", back_populates="notification_type")
