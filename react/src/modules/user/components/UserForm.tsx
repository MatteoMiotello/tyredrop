import React from "react";
import {FormHandler} from "../../../common/components/shelly-ui/Form/useForm";
import {User} from "../../../__generated__/graphql";
import {Form, Input} from "../../../common/components/shelly-ui";
import {isEmail} from "../../../common/components/shelly-ui/Input";
import {isRequired} from "../../../common/validation/validators";

type UserFormProps = {
    form: FormHandler
    user?: User
}
const UserForm: React.FC<UserFormProps> = ( {form, user} ) => {
    return <Form form={form} saveForm={() => true}>
        <Form.GridLayout>
        <Input.FormControl>
            <Input.Label>
                Email
            </Input.Label>
            <Input type="email" {...form.registerInput({
                name: 'email',
                validators: [isEmail( "Il campo deve essere una email valida" )],
                disable: Boolean( user )
            })}/>
        </Input.FormControl>
        <Input.FormControl className="col-span-6">
            <Input.Label>
                Nome
            </Input.Label>
            <Input {...form.registerInput({
                name: 'name',
                validators: [isRequired( "nome" )]
            })}/>
        </Input.FormControl>
            <Input.FormControl className="col-span-6">
            <Input.Label>
                Cognome
            </Input.Label>
            <Input {...form.registerInput({
                name: 'surname',
                validators: [isRequired( "cognome" )]
            })}/>
        </Input.FormControl>
        </Form.GridLayout>
    </Form>;
};

export default UserForm;