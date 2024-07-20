import { useState } from "react";

export default function App() {
  const [sent, setSent] = useState<string>("Nothing sent yet...");
  const [received, setReceived] = useState<string>("Nothing received yet...");
  const [outputs, setOutputs] = useState<string[]>([]);

  const makeRequest = async (
    url: string,
    payload: object,
    method: string,
    service: string
  ) => {
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
    setOutputs([
      `response from ${service}`,
      data.message,
      new Date().toString(),
    ]);
  };

  const OrderCreation = async () => {
    const payload = {
      name: "jateen",
      email: "kalsijatin67@icloud.com",
      phone: "4389893868",
      course_id: 1,
    };

    await makeRequest(
      "http://localhost:5173/order",
      payload,
      "POST",
      "order creation"
    );
  };

  const twilioOrderEnable = async () => {
    const payload = {
      phone: "6789998212",
      course_id: 1,
      is_active: 1,
    };

    await makeRequest(
      "http://localhost:5173/order",
      payload,
      "PUT",
      "twilio order edit"
    );
  };

  const twilioOrderDisable = async () => {
    const payload = {
      phone: "6789998212",
      course_id: 1,
      is_active: 0,
    };

    await makeRequest(
      "http://localhost:5173/order",
      payload,
      "PUT",
      "twilio order edit"
    );
  };

  return (
    <div>
      <div>
        <div>
          <h1>open course</h1>
          <hr></hr>
          <a onClick={OrderCreation}>test order creation</a>
          <div>
            {outputs.length === 0 ? (
              <>
                <span>Output shows here...</span>
              </>
            ) : (
              <>
                <strong>Started</strong>
                <br></br>
                <i>Sending request...</i>
                <br></br>
                <strong>{outputs[0]}: </strong>
                <span>{outputs[1]}</span>
                <br></br>
                <strong>Ended: </strong>
                <span>{outputs[2]}</span>
              </>
            )}
          </div>
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
