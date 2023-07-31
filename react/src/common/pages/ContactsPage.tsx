import React from "react";
import Panel from "../components-library/Panel";
import ContactForm from "../components/ContactForm";
import {useForm} from "../components/shelly-ui";

const ContactsPage: React.FC = () => {
    const form = useForm();

    return <main className="md:p-24 p-4 flex items-center">
        <Panel className="w-full">
            <Panel.Title>
                Contatti
            </Panel.Title>
            <div className="grid grid-cols-2">
                <div className="ml-2">
                    <h2 className="text-2xl font-semibold mb-2">SDR TRADING SRLS</h2>
                    <p className="mb-2"><strong>Indirizzo:</strong> Via Roma, 220</p>
                    <p className="mb-2"><strong>CAP:</strong> 35020</p>
                    <p className="mb-2"><strong>Località:</strong> Albignasego (PD)</p>
                    <p className="mb-2"><strong>Tel.:</strong> 049 6898916</p>
                </div>
                <div>
                    Contatto
                    <ContactForm/>
                </div>
            </div>
        </Panel>
    </main>;
};

export default ContactsPage;