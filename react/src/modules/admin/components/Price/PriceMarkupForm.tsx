import React, {useEffect} from "react";
import {ProductPriceMarkup} from "../../../../__generated__/graphql";
import {useMutation} from "../../../../common/backend/graph/hooks";
import {UPDATE_MARKUP} from "../../../../common/backend/graph/mutation/price";
import {Form, Input} from "../../../../common/components/shelly-ui";
import {FormHandler} from "../../../../common/components/shelly-ui/Form/useForm";
import {isRequired} from "../../../../common/components/shelly-ui/Input";
import Spinner from "../../../../common/components/Spinner";

type PriceMarkupFormProps = {
    form: FormHandler
    markup?: ProductPriceMarkup
}
const PriceMarkupForm: React.FC<PriceMarkupFormProps> = ({form, markup}) => {
    const [mutate, queryMutation] = useMutation(UPDATE_MARKUP);

    useEffect(() => {
        if (!markup) {
            return;
        }

        form.setFormValues({
            markup: markup.markupPercentage
        });
    }, [markup]);

    return <div className="relative">
            {
                queryMutation.loading && <Spinner></Spinner>
            }
        <Form form={form}
              saveForm={(data) => mutate({
                  variables: {
                      id: markup?.id as string,
                      markupPercentage: data.markup
                  }
              })}>
            <Input.FormControl>
                <Input.Label>
                    Percentuale di ricarico
                </Input.Label>
                <Input {...form.registerInput({
                    name: 'markup',
                    validators: [isRequired('La percentuale e` richiesta')]
                })}/>
            </Input.FormControl>
        </Form>
    </div>;
};
export default PriceMarkupForm;