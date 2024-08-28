import { Nav } from "../../components/Nav/Nav";
import { useState, useEffect } from "react";
import Colors, { applyTheme } from "../../styles/ColorSystem";
import MainContentStyles from "../../components/MainContent/MainContent.module.css";

export const AboutPage: React.FC = () => {
  const [currentTheme] = useState<keyof typeof Colors>("burgundy");

  useEffect(() => {
    applyTheme(Colors[currentTheme]);
  }, [currentTheme]);

  return (
    <>
      <Nav />
      <div className={MainContentStyles.AboutContent}>
        <h3>How does ConUAlerts work?</h3>
        <p>
          We constantly grab information from VSB Concordia to detect if a
          course seat or waitlist position has opened. If it has, you receive a
          text message.
        </p>
        <h3>
          I made an order for a course last semester, but that order has
          expired. Why?
        </h3>
        <p>
          If you signed up for course notifications for a certain semester, like
          Fall 2024, you must sign up for notifications for that same course in
          the following semester.
        </p>
        <h3>
          I received an alert that a course has opened, but there weren't any
          when I checked. Why?
        </h3>
        <p>
          Someone sniped the spot before you could :( No worries, simply text
          "ORDERS" to view the ID of the course, then "START &lt;course
          number&gt;" to reinstate your order for that course!
        </p>
        <h3>How can I contact you?</h3>
        <p>
          For more details about our service, or any general information
          regarding ConUAlerts, feel free to reach out to us at{" "}
          <a href="mailto:help@conualerts.com">help@conualerts.com</a>. We are
          here to provide support and answer any questions you may have.
        </p>
        <br></br>
        <p>Made with &#9829; by Rei Kong, Danny Mousa, Jatin Kalsi</p>
      </div>
    </>
  );
};
