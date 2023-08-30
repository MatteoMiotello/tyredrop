import React, {useState} from "react";
import {CreateInvoiceMutation, CreateInvoiceMutationVariables} from "../../../../__generated__/graphql";
import {useMutation} from "../../../../common/backend/graph/hooks";
import {CREATE_INVOICE} from "../../../../common/backend/graph/mutation/invoice";
import {Form, Input} from "../../../../common/components/shelly-ui";
import {FormHandler} from "../../../../common/components/shelly-ui/Form/useForm";
import {useToast} from "../../../../store/toast";
import UserBillingSelect from "../User/UserBillingSelect";

type InvoiceAdminFormProps = {
    form: FormHandler
}
const InvoiceAdminForm: React.FC<InvoiceAdminFormProps> = ( {form} ) => {
    const [file, setFile] = useState<File | undefined>(undefined);
    const toastr = useToast();
    const [mutate] = useMutation<CreateInvoiceMutation, CreateInvoiceMutationVariables>(CREATE_INVOICE);


    return <Form form={form} saveForm={(data) => {
        if (file === undefined) {
            toastr.error( "Nessun file selezionato" );
            return false;
        }

        if ( !data.billing ) {
            toastr.error( "Nessun intestatario selezionato" );
            return false;
        }

        return mutate({
            variables: {
                number: data.number,
                billingId: data.billing,
                file: file
            }
        });
    }}>
        <Input.FormControl>
            <Input.Label>
                Numero
            </Input.Label>
            <Input {...form.registerInput({name: 'number'})}/>
        </Input.FormControl>
        <Input.FormControl>
            <Input.Label>
                Intestatario
            </Input.Label>
            <UserBillingSelect name="billing"/>
        </Input.FormControl>
        <Input.FormControl>
        <Input.Label>
            File
        </Input.Label>
        <input
            type="file"
            accept=".pdf,.docx,.pdfa"
            className="file-input"
            onChange={ (e) => {
                if ( !e.target.files ) {
                    return;
                }

                const file = e.target.files.item(0);

                if ( !file ) {
                    return;
                }

                if (file && file.size > 64000000) {
                    toastr.error( "Il file selezionato e` troppo grande" );
                    return;
                }

                setFile(file);
            } }/>
        </Input.FormControl>
    </Form> ;
};

export default InvoiceAdminForm;