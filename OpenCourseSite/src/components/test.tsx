import { useState } from "react";

export default function Test() {
  const [sent, setSent] = useState<string>("Nothing sent yet...");
  const [received, setReceived] = useState<string>("Nothing received yet...");

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
      name: "jateen",
      email: "kalsijatin67@icloud.com",
      phone: "4389893868",
      course_id: 1,
    };

    await makeRequest("http://localhost:8081/", payload, "POST");
  };

  const twilioOrderEnable = async () => {
    const payload = {
      phone: "6789998212",
      course_id: 1,
      is_active: 1,
    };

    await makeRequest("http://localhost:8081/", payload, "PUT");
  };

  const twilioOrderDisable = async () => {
    const payload = {
      phone: "6789998212",
      course_id: 1,
      is_active: 0,
    };

    await makeRequest("http://localhost:8081/", payload, "PUT");
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
