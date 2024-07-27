import { ReactNode, forwardRef } from "react"
import * as Select from "@radix-ui/react-select";
import { ChevronDownIcon, ChevronUpIcon } from "@radix-ui/react-icons";
import Colors from "../../styles/ColorSystem"
import SelectStyles from "./SelectMenu.module.css"

interface SelectMenuProps {
  setCurrentTheme: (theme: keyof typeof Colors) => void;
}

export const SelectMenu = (props: SelectMenuProps) => {
  return (
    <Select.Root>
        <Select.Trigger className={SelectStyles.Trigger} aria-label="University Menu">
        <Select.Value placeholder="University" />
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
                  <SelectItem value="mcgill-university" onClick={() => props.setCurrentTheme("blue")}>McGill University</SelectItem>
                  <SelectItem value="concordia-university" onClick={() => {
                    console.log('Concordia clicked');
                    props.setCurrentTheme('burgundy');
                  }}>Concordia University</SelectItem>
                  <SelectItem value="unt" onClick={() => props.setCurrentTheme("green")}>University of Northern Texas</SelectItem>
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

export interface SelectItemProps {
  value: string;
  children: ReactNode;
  className: string;
  asChild?: boolean;
  onClick?: () => void;
};

const SelectItemText = forwardRef<HTMLDivElement, SelectItemProps>(({ children, className, onClick, ...props }, forwardedRef) => {
  return (
    <div className={`${SelectStyles.Item} ${className}`} {...props} ref={forwardedRef} onClick={() => {
      console.log('CustomSelectItem clicked');
      if (onClick) onClick();
    }}>
      <Select.ItemText>{children}</Select.ItemText>
    </div>
  );
});

const SelectItem = forwardRef<HTMLDivElement, SelectItemProps>(({ children, className, onClick, value, ...props }, forwardedRef) => {
  return (
    <Select.Item asChild {...props} value={value} ref={forwardedRef}>
      <SelectItemText className={className} onClick={onClick}>
        {children}
      </SelectItemText>
    </Select.Item>
  );
});