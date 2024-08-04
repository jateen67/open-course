import * as RadioGroup from '@radix-ui/react-radio-group';
import FormStyles from "./Form.module.css"
import RadioGroupStyles from "./RadioGroup.module.css";
import CourseCombobox from './Combobox';
import CheckboxGroup from './CheckboxGroup';
import Checkbox from './Checkbox';

const Form = () => (
    <form className={FormStyles.Content}>
        <div className={FormStyles.SectionContent}>
            <h3>Term</h3>
            <RadioGroup.Root className={RadioGroupStyles.Root} defaultValue="default" aria-label="View density">
                <div className={RadioGroupStyles.Container}>
                    <RadioGroup.Item className={RadioGroupStyles.Item} value="default" id="r1">
                    <RadioGroup.Indicator className={RadioGroupStyles.Indicator} />
                    </RadioGroup.Item>
                    <label className={RadioGroupStyles.Label} htmlFor="r1">
                        Default
                    </label>
                </div>
                <div className={RadioGroupStyles.Container}>
                    <RadioGroup.Item className={RadioGroupStyles.Item} value="comfortable" id="r2">
                    <RadioGroup.Indicator className={RadioGroupStyles.Indicator} />
                    </RadioGroup.Item>
                    <label className={RadioGroupStyles.Label} htmlFor="r2">
                        Comfortable
                    </label>
                </div>
                <div className={RadioGroupStyles.Container}>
                <RadioGroup.Item className={RadioGroupStyles.Item} value="meow" id="r3">
                    <RadioGroup.Indicator className={RadioGroupStyles.Indicator} />
                    </RadioGroup.Item>
                    <label className={RadioGroupStyles.Label} htmlFor="r3">
                        Meow
                    </label>
                </div>
            </RadioGroup.Root>
        </div>
        
        <div className={FormStyles.SectionContent}>
            <h3>Course</h3>
            <CourseCombobox />
        </div>
        <div className={FormStyles.SectionContent}>
            <h3>Section</h3>
            <CheckboxGroup />
            <Checkbox />
        </div>
        <div className={FormStyles.SectionContent}>
            <h3>Contact Info</h3>
            <fieldset className={FormStyles.Fieldset}>
                <label className={FormStyles.Label} htmlFor="name">Name</label>
                <input className={FormStyles.Input} id="name" />
            </fieldset>
            <fieldset className={FormStyles.Fieldset}>
                <label className={FormStyles.Label} htmlFor="email">Email</label>
                <input className={FormStyles.Input} id="email" />
            </fieldset>
            <fieldset className={FormStyles.Fieldset}>
                <label className={FormStyles.Label} htmlFor="pnum">Phone Number</label>
                <input className={FormStyles.Input} id="pnum" />
            </fieldset>
            <button className={FormStyles.Button}>Checkout</button>
        </div>
    </form>
);

export default Form;