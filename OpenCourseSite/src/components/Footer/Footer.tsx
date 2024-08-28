import FooterStyles from "./Footer.module.css";

export const Footer = () => {
  return (
    <div className={FooterStyles.Footer}>
      <p>&copy; ConUAlerts {new Date().getFullYear()}</p>
    </div>
  );
};
