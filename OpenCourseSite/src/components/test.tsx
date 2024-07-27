import { useCallback, useState } from "react";
import { faker } from "@faker-js/faker";
import { CourseService, OrderService } from "services";
import { Order } from "models";

export default function Test() {
  //not optimal to use many states, better to have 1 state for all but its a test component idc
  const [sent, setSent] = useState<string>("nothing sent yet...");
  const [received, setReceived] = useState<string>("nothing received yet...");
  const [numCourses, setNumCourses] = useState<number>(0);
  const [id, setId] = useState<string>();
  const [courseId, setcourseId] = useState<string>();
  const [email, setEmail] = useState<string>();

  const getAllCourses = () => {
    CourseService.GetAll().subscribe({
      next: (courses) => {
        console.log(courses);
        setNumCourses(courses.length);
      },
      error: (err) => console.log(err),
    });
  };

  const OrderCreation = () => {
    const payload: Order = new Order({
      name: faker.person.fullName(),
      email: faker.internet.email(),
      phone: faker.phone.number(),
      courseId: Math.floor(Math.random() * numCourses) + 1,
    });

    OrderService.CreateOrder(payload).subscribe({
      next: (newOrder) => {
        console.log(newOrder);
        setSent(JSON.stringify(payload, undefined, 4));
        setReceived(JSON.stringify(newOrder, undefined, 4));
      },
      error: (err) => console.log(err),
    });
  };

  const OrderUpdate = () => {
    const payload: Order = new Order({
      name: "FAKE NAME",
      email: "FAKE@EMIL.COM",
      phone: faker.phone.number(),
      courseId: Math.floor(Math.random() * numCourses) + 1,
    });

    OrderService.UpodateOrder(payload).subscribe({
      next: (updatedOrder) => {
        console.log(updatedOrder);
        setSent(JSON.stringify(payload, undefined, 4));
        setReceived(JSON.stringify(updatedOrder, undefined, 4));
      },
    });
  };

  const handleIdSubmit = () => {
    if (!id || isNaN(Number(id)) || Number(id) < 0) return;

    OrderService.GetOrderById(Number(id)).subscribe({
      next: (order) => setReceived(JSON.stringify(order, undefined, 4)),
      error: (err) => console.log(err),
    });
  };

  const handleEmailSubmit = () => {
    if (!email) return;

    OrderService.GetOrderByEmail(email).subscribe({
      next: (order) => setReceived(JSON.stringify(order, undefined, 4)),
      error: (err) => console.log(err),
    });
  };

  const handleCourseIdSubmit = () => {
    if (!courseId || isNaN(Number(courseId)) || Number(courseId) < 0) return;

    OrderService.GetOrderByCourseId(Number(courseId)).subscribe({
      next: (order) => setReceived(JSON.stringify(order, undefined, 4)),
      error: (err) => console.log(err),
    });
  };

  return (
    <div>
      <div>
        <div>
          <h1>open course</h1>
          <hr></hr>
          <a onClick={getAllCourses}>test get all courses</a>
          <br></br>
          <a onClick={OrderCreation}>test order creation (hardcoded value)</a>
          <br></br>
          <a onClick={OrderUpdate}>test order update (hardcoded value)</a>
          <br></br>
          <a onClick={OrderUpdate}>test get all orders</a>
          <br></br>
          <div>
            <p>test get order by Id</p>
            <input
              type="text"
              value={id}
              onChange={(event) => setId(event?.target.value)}
              placeholder="Enter an ID"
            />
            <button onClick={handleIdSubmit}>Send Request</button>
          </div>
          <br></br>
          <div>
            <p>test get order by email</p>
            <input
              type="text"
              value={email}
              onChange={(event) => setEmail(event?.target.value)}
              placeholder="Enter an email"
            />
            <button onClick={handleEmailSubmit}>Send Request</button>
          </div>
          <br></br>
          <div>
            <p>test get order by courseId</p>
            <input
              type="text"
              value={courseId}
              onChange={(event) => setcourseId(event?.target.value)}
              placeholder="Enter a courseId"
            />
            <button onClick={handleCourseIdSubmit}>Send Request</button>
          </div>
          <br></br>
        </div>
      </div>
      <div>
        <div>
          <h4>Sent</h4>
          <div style={{ outline: "1px solid silver", padding: "2em" }}>
            <pre>
              <span style={{ fontWeight: "bold" }}>{sent}</span>
            </pre>
          </div>
        </div>
        <div>
          <h4>Received</h4>
          <div style={{ outline: "1px solid silver", padding: "2em" }}>
            <pre>
              <span style={{ fontWeight: "bold" }}>{received}</span>
            </pre>
          </div>
        </div>
        <div>
          <h4>Course num</h4>
          <div style={{ outline: "1px solid silver", padding: "2em" }}>
            <pre>
              <span style={{ fontWeight: "bold" }}>{numCourses}</span>
            </pre>
          </div>
        </div>
      </div>
    </div>
  );
}
