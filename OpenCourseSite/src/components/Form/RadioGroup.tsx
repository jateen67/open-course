import { Radio as HeadlessRadio, RadioGroup as HeadlessRadioGroup } from '@headlessui/react'
// import { CheckCircleIcon } from '@heroicons/react/24/solid'
import { useState } from "react";
import RadioGroupStyles from "./RadioGroup.module.css";

interface Option {
    id: string;
    value: string;
    label: string;
}

interface RadioGroupProps {
    options: Option[];
}

const RadioGroup: React.FC<RadioGroupProps> = ({ options }) =>  {
    const [selected, setSelected] = useState(options[0])

    return (
        <HeadlessRadioGroup
            value={selected}
            onChange={setSelected}
            aria-label="Term"
            className={RadioGroupStyles.RadioGroup}
        >
            {options.map((option) => (
                <HeadlessRadio
                    key={option.id}
                    value={option}
                    className={`${RadioGroupStyles.Radio} ${(selected.id === option.id) ? RadioGroupStyles.Checked : ""}`}
                >
                <div className={RadioGroupStyles.RadioContent}>
                    <div className={RadioGroupStyles.RadioLabel}>
                        <p>{option.label}</p>
                    </div>
                    {/* <CheckCircleIcon className={`${RadioGroupStyles.CheckCircle} ${(selected.id === option.id) ? RadioGroupStyles.Checked : ""}`} /> */}
                </div>
                </HeadlessRadio>
            ))}
        </HeadlessRadioGroup>
    );
}

export default RadioGroup;