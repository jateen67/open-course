export type SemesterOption = {
    id: number;
    label: string;
    value: "2242" | "2244" | "2243" | "2245";
}

export interface ThemeContextType {
    currentTheme: string;
    setCurrentTheme: (theme: string) => void;
}
