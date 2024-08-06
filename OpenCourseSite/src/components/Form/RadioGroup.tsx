import * as RadixRadio from '@radix-ui/react-radio-group';
import RadioGroupStyles from "./RadioGroup.module.css";

interface Option {
    id: string;
    value: string;
    label: string;
}

interface RadioGroupProps {
    options: Option[];
}

const RadioGroup: React.FC<RadioGroupProps> = ({ options }) => (
    <RadixRadio.Root
        className={RadioGroupStyles.Root}
        defaultValue="default"
        aria-label="View density"
    >
        {options.map((option) => (
            <div key={option.id} className={RadioGroupStyles.Container}>
                <RadixRadio.Item className={RadioGroupStyles.Item} value={option.value} id={option.id}>
                    <RadixRadio.Indicator className={RadioGroupStyles.Indicator} />
                </RadixRadio.Item>
                <label className={RadioGroupStyles.Label} htmlFor={option.id}>{option.label}</label>
            </div>
        ))}
    </RadixRadio.Root>
);

export default RadioGroup;