import React from "react";
import {useQuery} from "../../../common/backend/graph/hooks";
import {PRICE_MARKUPS} from "../../../common/backend/graph/query/products";
import Panel from "../../../common/components-library/Panel";
import {Button, Modal, useForm, useModal} from "../../../common/components/shelly-ui";
import PriceMarkupTable from "../components/Price/PriceMarkupTable";

const PriceMarkupPage: React.FC = () => {
    const query = useQuery(PRICE_MARKUPS);
    const modal = useModal();
    const form = useForm();
    return <main>
        <Modal modal={modal}>
            <Modal.Title>
                Modifica un prezzo
            </Modal.Title>
            <Modal.Actions>
                <Button onClick={modal.close}>
                    Annulla
                </Button>
                <Button onClick={form.submitForm} buttonType="primary">
                    Salva
                </Button>
            </Modal.Actions>
        </Modal>
        <Panel>
            <Panel.Title>
                Tutti i prezzi
            </Panel.Title>
            {
                query.data &&
                <PriceMarkupTable query={query}/>
            }
        </Panel>
    </main>;
};

export default PriceMarkupPage;