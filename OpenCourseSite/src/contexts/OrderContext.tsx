import { createContext, useContext, useState } from "react";

interface OrderContextProps {
    email: string;
    phone: string;
    classNumber: number | null;
    setEmail: (email: string) => void;
    setPhone: (phone: string) => void;
    setClassNumber: (classNumber: number | null) => void;
}

const OrderContext = createContext<OrderContextProps | undefined>(undefined);

export const useOrderContext = () => {
    const context = useContext(OrderContext);
    if (!context) {
        throw new Error("useOrderContext must be used within a OrderProvider");
    }
    return context;
};

export const OrderProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [email, setEmail] = useState<string>("");
    const [phone, setPhone] = useState<string>("");
    const [classNumber, setClassNumber] = useState<number | null>(null);

    return (
        <OrderContext.Provider value={{
            email,
            phone,
            classNumber,
            setEmail,
            setPhone,
            setClassNumber
        }}>
            {children}
        </OrderContext.Provider>
    );
};