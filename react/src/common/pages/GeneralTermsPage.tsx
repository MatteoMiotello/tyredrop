import React from "react";
import Panel from "../components-library/Panel";

const GeneralTermsPage: React.FC = () => {
    return <main className="p-4 md:p-24 flex items-center">
        <Panel className="w-full">
            <h1 className="text-4xl font-bold mb-4">Condizioni di Vendita valide fino al 31.12.2023.</h1>

            <div className="container mx-auto p-4">
                {/* SEZIONE ORDINI IN SPEDIZIONE */}
                <section className="mb-8">
                    <h2 className="text-2xl font-semibold mb-4">ORDINI IN SPEDIZIONE:</h2>
                    <p className="mb-4">Gli ordini partiranno da un minimo di 2 pezzi, con spedizione gratuita presso la Vostra sede (per il pezzo singolo addebiteremo €6,90 di spese di spedizione).</p>
                    <p className="mb-4">Per quanto concerne la spedizione nelle isole, Sicilia e Sardegna, o zone periferiche difficilmente raggiungibili ci sarà l’incremento di €15,00 per pacchetto. *(Il pacchetto viene inteso con due gomme fino a R18 compreso, per raggio superiore la gomma è singola in ogni pacchetto).</p>
                    <p>*eventuali CAP specifici con tali problematiche sono inseriti nella sezione Informazioni.</p>
                </section>

                {/* SEZIONE CONVENZIONATO INSTALLATORE B2C */}
                <section className="mb-8">
                    <h2 className="text-2xl font-semibold mb-4">CONVENZIONATO INSTALLATORE B2C</h2>
                    <p className="mb-4">Lavora con noi e diventa installatore. Noi ti mandiamo i clienti con il prezzo da noi convenzionato di €40 per 4 pneumatici più €40 assetto ruote (convergenza). Ti inseriamo nel nostro database come installatore con il tuo indirizzo dell’officina, il cliente paga le prestazioni in officina.</p>
                </section>

                {/* SEZIONE PERCHÈ ACQUISTARE DA NOI */}
                <section>
                    <h2 className="text-2xl font-semibold mb-4">Perchè acquistare da noi:</h2>
                    <ul className="list-disc ml-8">
                        <li className="mb-2">Prezzi sempre concorrenziali</li>
                        <li className="mb-2">Nessun canone mensile</li>
                        <li className="mb-2">Assistenza sempre pronta con le nostre 20 ragazze</li>
                        <li className="mb-2">Piattaforma esclusiva italiana</li>
                    </ul>
                </section>

            </div>
        </Panel>
    </main>;
};

export default GeneralTermsPage;