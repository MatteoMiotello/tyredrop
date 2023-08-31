import React, {useState} from "react";
import {Button, Form, Input, useForm} from "../../common/components/shelly-ui";
import Panel from "../../common/components-library/Panel";
import {createBackendClient} from "../../common/backend/backend-client";
import {useToast} from "../../store/toast";
import {isRequired} from "../../common/validation/validators";
import {isEmail} from "../../common/components/shelly-ui/Input";
import {Link} from "react-router-dom";

const ResetPasswordPage: React.FC = () => {
    const form = useForm();
    const [email, setEmail] = useState<string | null>();
    const toastr = useToast();
    const client = createBackendClient();

    const submit = (data: any) => {
        if (!data.email) {
            return false;
        }

        return client.resetPassword(data.email)
            .then(() => {
                setEmail(email);
                toastr.success("Email inviata con successo");
            })
            .catch(() => {
                setEmail(undefined);
                toastr.error("C'è stato un erore nel reset della password");
            });
    };

    return <>
        <div className="flex flex-col justify-center items-center my-auto ">
            <Panel className="w-1/2">
                {
                    email ?
                        <p>Una email è stata inviata a {email}</p>
                        :
                        <div>
                            <p className="my-4 text-lg">
                                Per procedere con il reset della password è necessario inserire l'email del tuo account.
                            </p>
                            <Form form={form} saveForm={submit}>
                                <Input.FormControl>
                                    <Input.Label>
                                        Email
                                    </Input.Label>
                                    <Input
                                        placeholder="email"
                                        type="email" {...form.registerInput({
                                        name: 'email',
                                        validators: [isRequired("email"), isEmail( "Email non valida" )]
                                    })}/>
                                    <Form.FormButtons>
                                        <Link to="/" className="btn">
                                            Torna al login
                                        </Link>
                                        <Button type="submit" buttonType="primary">
                                            Invia Email
                                        </Button>
                                    </Form.FormButtons>
                                </Input.FormControl>
                            </Form>
                        </div>
                }
            </Panel>
        </div>
    </>;
};

export default ResetPasswordPage;