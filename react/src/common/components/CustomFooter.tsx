import React from "react";
import Footer from "../components-library/Footer";
import {Link} from "react-router-dom";
import ContactForm from "./ContactForm";

const CustomFooter: React.FC = () => {
    return <Footer>
        <Footer.Column>
            <span className="text-lg font-semibold uppercase">Link utili</span>
            <Footer.Element>
                <Link to="/contacts">
                    Contatti
                </Link>
            </Footer.Element>
            <Footer.Element>
                <Link to="/general-terms">
                    Condizioni generali di vendita
                </Link>
            </Footer.Element>
        </Footer.Column>
        <Footer.Column>
        </Footer.Column>
        <Footer.Column className="w-full">
            <div className="ml-auto">
            <span className="text-lg font-semibold uppercase">Richiesta di contatto</span>
            <ContactForm/>
            </div>
        </Footer.Column>
    </Footer>;
};
export default CustomFooter;