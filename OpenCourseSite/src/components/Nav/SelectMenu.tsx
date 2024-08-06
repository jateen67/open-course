import { forwardRef } from "react"
import * as Select from "@radix-ui/react-select";
import { ChevronDownIcon, ChevronUpIcon } from "@radix-ui/react-icons";
import Colors from "../../styles/ColorSystem"
import SelectStyles from "./SelectMenu.module.css"

interface SelectMenuProps {
  setCurrentTheme: (theme: keyof typeof Colors) => void;
}

interface SelectItemProps extends React.ComponentPropsWithoutRef<"div"> {
  value: string;
  children: React.ReactNode;
};

const SelectItem = forwardRef<HTMLDivElement, SelectItemProps>(({ children, ...props }, forwardedRef) => {
  return (
    <Select.Item className={SelectStyles.Item} {...props} ref={forwardedRef}>
      <Select.ItemText>{children}</Select.ItemText>
    </Select.Item>
  );
});

export const SelectMenu: React.FC<SelectMenuProps> = ({ setCurrentTheme }) => {
  const handleSelect = (value: string) => {
    if (Object.keys(Colors).includes(value)) {
      setCurrentTheme(value as keyof typeof Colors);
    }
  };

  return (
    <Select.Root defaultValue="red" onValueChange={handleSelect}>
        <Select.Trigger className={SelectStyles.Trigger} aria-label="University Menu">
        <Select.Value />
        <Select.Icon className={SelectStyles.Icon}>
            <ChevronDownIcon />
        </Select.Icon>
        </Select.Trigger>
        <Select.Portal>
          <Select.Content className={SelectStyles.Content}>
              <Select.ScrollUpButton className={SelectStyles.ScrollButton}>
              <ChevronUpIcon />
              </Select.ScrollUpButton>
              <Select.Viewport className={SelectStyles.Viewport}>
                <Select.Group>
                  <SelectItem value="red">McGill University</SelectItem>
                  <SelectItem value="burgundy">Concordia University</SelectItem>
                  <SelectItem value="green">University of Northern Texas</SelectItem>
                </Select.Group>
              </Select.Viewport>
              <Select.ScrollDownButton className={SelectStyles.ScrollButton}>
                <ChevronDownIcon />
              </Select.ScrollDownButton>
          </Select.Content>
        </Select.Portal>
    </Select.Root>
  );
};