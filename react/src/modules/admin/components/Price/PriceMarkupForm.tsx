import React, {useEffect} from "react";
import {ProductPriceMarkup} from "../../../../__generated__/graphql";
import {useMutation} from "../../../../common/backend/graph/hooks";
import {CREATE_MARKUP, UPDATE_MARKUP} from "../../../../common/backend/graph/mutation/price";
import BrandField from "../../../../common/components/BrandField";
import {Form, Input} from "../../../../common/components/shelly-ui";
import {FormHandler} from "../../../../common/components/shelly-ui/Form/useForm";
import {isRequired} from "../../../../common/components/shelly-ui/Input";
import Spinner from "../../../../common/components/Spinner";
import SpecificationField from "../../../product/components/SpecificationField";
import {ProductSpecificationsSet} from "../../../product/enums/product-specifications-set";

type PriceMarkupFormProps = {
    form: FormHandler
    markup?: ProductPriceMarkup
}
const PriceMarkupForm: React.FC<PriceMarkupFormProps> = ({form, markup}) => {
    const [mutate, queryMutation] = useMutation(markup ? UPDATE_MARKUP : CREATE_MARKUP);

    useEffect(() => {
        if (!markup) {
            return;
        }

        form.setFormValues({
            markup: markup.markupPercentage,
            brand: markup.brand?.id,
            value: markup.productSpecificationValue?.id
        });
    }, [markup]);

    return <div className="relative">
        {
            queryMutation.loading && <Spinner></Spinner>
        }
        <Form form={form}
              saveForm={(data) =>{
                  return mutate({
                  variables: {
                      id: markup?.id as string,
                      input: {
                          brandId: data.brand,
                          specificationValueId: data.value,
                          markupPercentage: data.markup
                      }
                  }
              });}}>
            <Form.GridLayout>
                {
                    !markup && <>
                        <Input.FormControl className="col-span-6">
                            <Input.Label>
                                Brand
                            </Input.Label>
                            <BrandField toId={true} {...form.registerInput({name: 'brand', disable: Boolean(markup)})}/>
                        </Input.FormControl>
                        <Input.FormControl className="col-span-6">
                            <Input.Label>
                                Raggio
                            </Input.Label>
                            <SpecificationField
                                specificationCode={ProductSpecificationsSet.TYRE.RIM} {...form.registerInput({
                                name: "value",
                                disable: Boolean(markup)
                            })}
                                toId={true}/>
                        </Input.FormControl>
                    </>
                }
                <Input.FormControl className="col-span-12">
                    <Input.Label>
                        Percentuale di ricarico
                    </Input.Label>
                    <Input {...form.registerInput({
                        name: 'markup',
                        validators: [isRequired('La percentuale e` richiesta')]
                    })}/>
                </Input.FormControl>
            </Form.GridLayout>
        </Form>
    </div>;
};
export default PriceMarkupForm;