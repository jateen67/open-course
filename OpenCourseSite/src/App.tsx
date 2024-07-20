import { useState } from "react";

export default function App() {
  const [sent, setSent] = useState<string>("Nothing sent yet...");
  const [received, setReceived] = useState<string>("Nothing received yet...");
  const [outputs, setOutputs] = useState<string[]>([]);

  const makeRequest = async (url: string, payload: object, service: string) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const res = await fetch(url, {
      method: "POST",
      body: JSON.stringify(payload),
      headers: headers,
    });

    const data = await res.json();

    setSent(JSON.stringify(payload, undefined, 4));
    setReceived(JSON.stringify(data, undefined, 4));
    setOutputs([
      `Response from ${service}`,
      data.message,
      new Date().toString(),
    ]);
  };

  const getAuthentication = async () => {
    const payload = {
      email: "admin@example.com",
      password: "password123",
    };

    await makeRequest("http://localhost:5173/orders", payload, "Authenticator");
  };

  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <h1 className="mt-5 text-light">Go Distributed System</h1>
          <hr></hr>
          <a
            className="btn btn-outline-secondary text-light"
            onClick={getAuthentication}
          >
            Test Auth
          </a>
          <div
            className="mt-5"
            style={{ outline: "1px solid silver", padding: "2em" }}
          >
            {outputs.length === 0 ? (
              <>
                <span className="text-muted">Output shows here...</span>
              </>
            ) : (
              <>
                <strong className="text-success">Started</strong>
                <br></br>
                <i className="text-light">Sending request...</i>
                <br></br>
                <strong className="text-light">{outputs[0]}: </strong>
                <span className="text-light">{outputs[1]}</span>
                <br></br>
                <strong className="text-danger">Ended: </strong>
                <span className="text-light">{outputs[2]}</span>
              </>
            )}
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col">
          <h4 className="mt-5 text-light">Sent</h4>
          <div
            className="mt-1"
            style={{ outline: "1px solid silver", padding: "2em" }}
          >
            <pre>
              <span className="text-light" style={{ fontWeight: "bold" }}>
                {sent}
              </span>
            </pre>
          </div>
        </div>
        <div className="col">
          <h4 className="mt-5 text-light">Received</h4>
          <div
            className="mt-1"
            style={{ outline: "1px solid silver", padding: "2em" }}
          >
            <pre>
              <span className="text-light" style={{ fontWeight: "bold" }}>
                {received}
              </span>
            </pre>
          </div>
        </div>
      </div>
    </div>
  );
}
