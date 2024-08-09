import Checkbox from "./Checkbox";
import CheckboxGroupStyles from "./CheckboxGroup.module.css"

interface Section {
    sectionType: string;
    sectionNum: string;
    crn: number;
    days: string[];
    times: string[];
}

interface CheckboxGroupProps {
    sections: Section[];
}

const headers = [
    { label: "Section", containerClass: CheckboxGroupStyles.SectionContainer },
    { label: "CRN", containerClass: CheckboxGroupStyles.CRNContainer },
    { label: "Day", containerClass: CheckboxGroupStyles.DayContainer },
    { label: "Time", containerClass: CheckboxGroupStyles.TimeContainer }
];

const CheckboxGroup: React.FC<CheckboxGroupProps> = ({ sections }) => (
    <div className={CheckboxGroupStyles.Container}>
        <div className={CheckboxGroupStyles.HeaderRow}>
            <div className={CheckboxGroupStyles.BoxContainer}>
            </div>
            {headers.map(header => (
                <div className={header.containerClass} key={header.label}>
                    <p className={CheckboxGroupStyles.Heading}>{header.label}</p>
                </div>
            ))}
        </div>
        {sections.map((section) => (
            <Checkbox
                section={section.sectionType + " " + section.sectionNum}
                crn={section.crn}
                days={section.days}
                times={section.times}
            />
        ))}
    </div>
);

export default CheckboxGroup;