import { Button as HeadlessButton } from "@headlessui/react";
import ButtonStyles from "./Button.module.css";
import { useFormContext, useOrderContext } from "../../contexts";
import { OrderService } from "services";
import { Order } from "models";

export const Button: React.FC = () => {
  const { phone } = useOrderContext();
  const { selectedCheckboxes } = useFormContext();

  const createOrder = (classNumber: number) => {
    const newOrder = new Order({ classNumber, phone });
    OrderService.CreateOrder(newOrder).subscribe({
      next: (order) => console.log("Order created successfully:", order),
      error: (error) => console.error("Error creating order:", error),
    });
  };

  const handleSubmit = () => {
    if (!phone) {
      return;
    }
    for (const classNumber of selectedCheckboxes) {
      console.log(classNumber);
      createOrder(classNumber);
    }
  };

  return (
    <HeadlessButton
      className={ButtonStyles.Button}
      type="submit"
      onClick={handleSubmit}
      disabled={!phone}
    >
      Checkout
    </HeadlessButton>
  );
};
