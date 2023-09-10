import React, {useState} from "react";
import {Button, Form, Input, useForm} from "../../common/components/shelly-ui";
import Panel from "../../common/components-library/Panel";
import {createBackendClient} from "../../common/backend/backend-client";
import {useToast} from "../../store/toast";
import {isRequired} from "../../common/validation/validators";
import {isEmail} from "../../common/components/shelly-ui/Input";
import {Link} from "react-router-dom";
import {Simulate} from "react-dom/test-utils";
import load = Simulate.load;
import Spinner from "../../common/components/Spinner";

const ResetPasswordPage: React.FC = () => {
    const form = useForm();
    const [email, setEmail] = useState<string | null>();
    const [loading, setLoading] = useState(false);
    const toastr = useToast();
    const client = createBackendClient();

    const submit = (data: any) => {
        if (!data.email) {
            return false;
        }

        setLoading(true);
        return client.resetPassword(data.email)
            .then(() => {
                setEmail(data.email);
                toastr.success("Email inviata con successo");
            })
            .catch((err) => {
                setEmail(undefined);
                if ( err.response.data.status_code == 4001 ) {
                    toastr.error( "Utente non trovato" );
                    return;
                }

                toastr.error("C'è stato un erore nel reset della password");
            }).finally( () =>setLoading(false) );
    };

    return <>
        <div className="flex flex-col justify-center items-center my-auto ">
            {loading && <Spinner></Spinner>}
            <Panel className="w-full md:w-1/2">
                {
                    email ?
                        <p>Una email è stata inviata a {email} <Link className="link-accent" to="/auth/login"> Torna al login </Link></p>
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