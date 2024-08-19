export type SemesterOption = {
    id: "2242" | "2244" | "2243" | "2245";
    label: string;
}

export interface ThemeContextType {
    currentTheme: string;
    setCurrentTheme: (theme: string) => void;
}