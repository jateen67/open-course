const commonColors = {
    primary900: "rgb(24, 24, 27)",
    primary800: "rgb(39, 39, 42)",
    primary700: "rgb(63, 63, 70)",
    primary650: "rgb(74, 74, 83)",
    primary600: "rgb(82, 82, 91)",
    primary500: "rgb(113, 113, 122)",
    primary400: "rgb(161, 161, 170)",
    primary300: "rgb(212, 212, 216)",
    secondary900: "rgb(17, 24, 39)",
    secondary800: "rgb(31, 41, 55)",
    secondary700: "rgb(55, 65, 81)",
    secondary650: "rgb(67, 77, 94)",
    secondary600: "rgb(75, 85, 99)",
    secondary500: "rgb(107, 114, 128)",
    secondary400: "rgb(156, 163, 175)",
    secondary300: "rgb(209, 213, 219)",
    secondary200: "rgb(229, 231, 235)",
    secondary100: "rgb(243, 244, 246)",
    brightBlue: "rgb(28, 100, 242)"
};

type CommonColors = typeof commonColors;

export type ThemeColors = {
    tertiary: string;
} & CommonColors;

const createTheme = (tertiary: string): ThemeColors => ({
    tertiary,
    ...commonColors,
});

const Colors = {
    burgundy: createTheme("rgb(40, 27, 29)"),
    blue: createTheme("rgb(34, 43, 56)"),
    green: createTheme("rgb(28, 36, 29)"),
};

export const applyTheme = (theme: ThemeColors) => {
    const root = document.documentElement;
    Object.keys(theme).forEach((key) => {
        root.style.setProperty(`--${key}`, theme[key as keyof ThemeColors]);
    });
};

export default Colors;