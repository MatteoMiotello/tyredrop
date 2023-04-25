import React, {PropsWithChildren} from "react";

interface PanelProps extends PropsWithChildren {
    className?: string
}

const Panel : React.FC<PanelProps> = ( props: PanelProps ) => {
    return <div className={"rounded-sm bg-secondary-content p-4 relative " + props.className}>
        { props.children }
    </div>;
};

export default Panel;