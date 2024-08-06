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

const CheckboxGroup: React.FC<CheckboxGroupProps> = ({ sections }) => (
    <div className={CheckboxGroupStyles.Container}>
        <div className={CheckboxGroupStyles.HeaderRow}>
            <div className={CheckboxGroupStyles.BoxContainer}>
            </div>
            <div className={CheckboxGroupStyles.SectionContainer}>
                <p>Section</p>
            </div>
            <div className={CheckboxGroupStyles.CRNContainer}>
                <p>CRN</p>
            </div>
            <div className={CheckboxGroupStyles.DayContainer}>
                <p>Day</p>
            </div>
            <div className={CheckboxGroupStyles.TimeContainer}>
                <p>Time</p>
            </div>
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