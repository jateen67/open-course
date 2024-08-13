export interface Course {
    id: number;
    courseCode: string;
    courseTitle: string;
    semester: string;
    days: string[];
    startTimeEndTime: string[];
    credits: string;
    section: string;
    openSeats: number;
    waitlistAvailable: number;
    waitlistCapacity: number;
    createdAt: string;
    updatedAt: string;
}

export interface Order {
    id: number;
    name: string;
    email: string;
    phone: string;
    courseId: number;
    isActive: boolean;
    createdAt: Date;
    updatedAt: Date
}

export type SemesterOption = {
    id: number;
    label: string;
    value: "fall" | "winter" | "fall|winter" | "summer";
}

export interface ThemeContextType {
    currentTheme: string;
    setCurrentTheme: (theme: string) => void;
}
