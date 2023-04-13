import React from "react";
import Footer, {FooterColumn} from "./Footer";
import {useTranslation} from "react-i18next";

const CustomFooter: React.FC = () => {
    const {t} = useTranslation()

    const footerData: FooterColumn[] = [
        {
            links: [
                {
                    title: t('footer.who'),
                },
                {
                    title: t('footer.newsletter'),
                },
                {
                    title: t('footer.assistance'),
                },
                {
                    title: t('footer.faq'),
                }
            ]
        }
    ]

    return <Footer data={footerData}></Footer>
}
export default CustomFooter