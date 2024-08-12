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
    createdAt: Date;
    updatedAt: Date;
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

export interface ThemeContextType {
    currentTheme: string;
    setCurrentTheme: (theme: string) => void;
}
