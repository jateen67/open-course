import { Nav } from "../../components/Nav/Nav";
import { Footer } from "components/Footer";
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
          text message. Your order then becomes disabled - you may renable it by
          texting "ORDERS" to view the ID of the course, then "START &lt;course
          ID&gt;" to reinstate the order.
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
          spots when I checked. Why?
        </h3>
        <p>
          Someone sniped the spot before you could :( No worries, simply
          reinstate the order as described above!
        </p>
        <p>
          It could also be because you are monitoring a section during time of
          the semester where certain seats are reserved for certain majors (e.g
          PSYC majors only). Because of this, you received a false positive
          since it says that there's an available spot on VSB, and it's not
          possible to account for this edge case. Once the reserved seats
          timeline ends, you should be eligible for any open spots you receive
          notifications for.
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
      <Footer />
    </>
  );
};
