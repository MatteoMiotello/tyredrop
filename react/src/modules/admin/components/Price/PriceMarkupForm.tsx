import React from "react";
import {Form} from "../../../../common/components/shelly-ui";
import {FormHandler} from "../../../../common/components/shelly-ui/Form/useForm";

type PriceMarkupFormProps = {
    form: FormHandler
}
const PriceMarkupForm: React.FC<PriceMarkupFormProps> = ({form}) => {
    return <Form form={form} saveForm={(data) => true}>

    </Form>;
};
export default PriceMarkupForm;