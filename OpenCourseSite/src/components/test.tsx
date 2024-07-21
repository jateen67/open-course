import { useEffect, useState } from "react";
import { faker } from "@faker-js/faker";

export default function Test() {
  const [sent, setSent] = useState<string>("nothing sent yet...");
  const [received, setReceived] = useState<string>("nothing received yet...");
  const [numCourses, setNumCourses] = useState<number>(0);

  useEffect(() => {
    getAllCourses();
  }, []);

  const getAllCourses = async () => {
    const response = await fetch("http://localhost:8081/courses");
    const data = await response.json();
    setNumCourses(data.length);
  };

  const makeRequest = async (url: string, payload: object, method: string) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const res = await fetch(url, {
      method: method,
      body: JSON.stringify(payload),
      headers: headers,
    });

    const data = await res.json();

    setSent(JSON.stringify(payload, undefined, 4));
    setReceived(JSON.stringify(data, undefined, 4));
  };

  const OrderCreation = async () => {
    const payload = {
      name: faker.person.fullName(),
      email: faker.internet.email(),
      phone: faker.phone.number(),
      course_id: Math.floor(Math.random() * numCourses) + 1,
    };

    await makeRequest("http://localhost:8081/orders", payload, "POST");
  };

  const twilioOrderEnable = async () => {
    const payload = {
      phone: "",
      course_id: -1,
      is_active: true,
    };

    await makeRequest("http://localhost:8081/orders", payload, "PUT");
  };

  const twilioOrderDisable = async () => {
    const payload = {
      phone: "",
      course_id: -1,
      is_active: false,
    };

    await makeRequest("http://localhost:8081/orders", payload, "PUT");
  };

  const mailerTest = async () => {
    const payload = {
      mail: {
        from: "me@example.com",
        to: "you@example.com",
        subject: "Test Email Subject",
        message: "Hello, world! This is my email",
      },
    };

    await makeRequest("http://localhost:8082/mail", payload, "POST");
  };

  return (
    <div>
      <div>
        <div>
          <h1>open course</h1>
          <hr></hr>
          <a onClick={OrderCreation}>test order creation</a>
          <br></br>
          <a onClick={twilioOrderEnable}>test twilio order enable</a>
          <br></br>
          <a onClick={twilioOrderDisable}>test twilio order disable</a>
          <br></br>
          <a onClick={getAllCourses}>test get all courses</a>
          <br></br>
          <a onClick={mailerTest}>test mailer</a>
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
      </div>
    </div>
  );
}
