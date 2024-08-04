import Checkbox from "./Checkbox";
import CheckboxGroupStyles from "./CheckboxGroup.module.css"

const CheckboxGroup = () => {
    return (
        <div className={CheckboxGroupStyles.Container}>
            <div className={CheckboxGroupStyles.Row}>
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
            <hr/>
            <Checkbox />
        </div>
    )
}

export default CheckboxGroup;