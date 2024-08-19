import { Button as HeadlessButton } from "@headlessui/react";
import ButtonStyles from "./Button.module.css";
import { useFormContext, useOrderContext } from "../../contexts";

export const Button: React.FC = () => {
    const { email, phone } = useOrderContext();
    const { selectedCheckboxes } = useFormContext();

    console.log("Button.tsx:" + email);
    console.log("Button.tsx:" + phone);
    console.log("Button.tsx:" + selectedCheckboxes);

    const createOrder = async (classNumber: number) => {
        try {
            const response = await fetch("http://localhost:8081/orders", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ classNumber, email, phone }),
            });

            if (!response.ok) {
                throw new Error("Failed to create order");
            }
    
            const data = await response.json();
            console.log("Order created successfully:", data);
        } catch (error) {
            console.error("Error creating order:", error);
        }
    
    };

    const handleSubmit = async () => {
        for (const classNumber of selectedCheckboxes) {
            console.log(classNumber);
            await createOrder(classNumber);
        }
    };

    return (
        <HeadlessButton className={ButtonStyles.Button} type="submit" onClick={handleSubmit}>
            Checkout
        </HeadlessButton>
    );
}