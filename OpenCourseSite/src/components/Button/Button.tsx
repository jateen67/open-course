import { Button as HeadlessButton } from "@headlessui/react";
import ButtonStyles from "./Button.module.css";
import { useFormContext, useOrderContext } from "../../contexts";
import { OrderService } from "services";
import { Order } from "models";
import { useSnackbar } from "notistack";

export const Button: React.FC = () => {
  const { phone } = useOrderContext();
  const { selectedCheckboxes } = useFormContext();
  const { enqueueSnackbar } = useSnackbar()

  const createOrder = (classNumber: number) => {
    const newOrder = new Order({ classNumber, phone });
    OrderService.CreateOrder(newOrder).subscribe({
      next: () => enqueueSnackbar("Order created successfuly", { variant: 'success' }),
      error: (error) => enqueueSnackbar(error.message, { variant: 'error' }),
    });
  };

  const handleSubmit = () => {
    if (!phone) {
      return;
    }
    for (const classNumber of selectedCheckboxes) {
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
