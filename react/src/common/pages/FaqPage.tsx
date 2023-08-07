import React from "react";
import Panel from "../components-library/Panel";
import MarkdownPage from "../components/MarkdownPage";

const FaqPage: React.FC = () => {
    return <Panel className="px-8">
        <MarkdownPage path="../../assets/pages/faq"/>
    </Panel>;
};

export default FaqPage;