import { ReactNode, forwardRef } from "react"
import * as Select from "@radix-ui/react-select";
import { ChevronDownIcon, ChevronUpIcon } from "@radix-ui/react-icons";
import Colors from "../../styles/ColorSystem"
import SelectStyles from "./SelectMenu.module.css"

interface SelectMenuProps {
  setCurrentTheme: (theme: keyof typeof Colors) => void;
}

export const SelectMenu: React.FC<SelectMenuProps> = ({ setCurrentTheme }) => {
  const handleSelect = (value: string) => {
    if (Object.keys(Colors).includes(value)) {
      setCurrentTheme(value as keyof typeof Colors);
    }
  };

  interface SelectItemProps extends React.ComponentPropsWithoutRef<'div'> {
    className?: string;
    value: string;
    children: React.ReactNode;
  };
  
  const SelectItem = forwardRef<HTMLDivElement, SelectItemProps>(({ children, className, ...props }, forwardedRef) => {
    return (
      <Select.Item className={`${SelectStyles.Item} ${className}`} {...props} ref={forwardedRef}>
        <Select.ItemText>{children}</Select.ItemText>
      </Select.Item>
    );
  });

  return (
    <Select.Root onValueChange={handleSelect}>
        <Select.Trigger className={SelectStyles.Trigger} aria-label="University Menu">
        <Select.Value placeholder="Select a University" />
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
                  <SelectItem value="blue">McGill University</SelectItem>
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