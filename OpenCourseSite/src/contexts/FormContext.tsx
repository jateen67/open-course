import React, { createContext, useContext, useState } from "react";
import { Course } from "../models";

interface FormContextProps {
    selectedTerm: string;
    query: string;
    selectedCourses: Course | null;
    selectedCheckboxes: number[];
    setSelectedTerm: (termCode: string) => void;
    setQuery: (query: string) => void;
    setSelectedCourses: (course: Course | null) => void;
    setSelectedCheckboxes: (courseIds: number[]) => void;
}

const FormContext = createContext<FormContextProps | undefined>(undefined);

export const useFormContext = () => {
    const context = useContext(FormContext);
    if (!context) {
        throw new Error("useFormContext must be used within a FormProvider");
    }
    return context;
};

export const FormProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [selectedTerm, setSelectedTerm] = useState<string>("");
    const [query, setQuery] = useState<string>("");
    const [selectedCourses, setSelectedCourses] = useState<Course | null>(null);
    const [selectedCheckboxes, setSelectedCheckboxes] = useState<number[]>([]);

    return (
        <FormContext.Provider value={{
            selectedTerm,
            query,
            selectedCourses,
            selectedCheckboxes,
            setSelectedTerm,
            setQuery,
            setSelectedCourses,
            setSelectedCheckboxes,
        }}>
            {children}
        </FormContext.Provider>
    );
};
