import classNames from "classnames";
import React, { ReactElement } from "react";
import Header from "../partials/Header";

export interface IUserPageContainerProps {
    onToggleDarkMode: () => void;
    isDarkMode: boolean;
    children: ReactElement<any, any>;
}

const UserPageContainer: React.FC<IUserPageContainerProps> = (props, { children }) => {

    return (
        <div className={classNames("h-screen dark:bg-dark", { 'dark': props.isDarkMode })}>
            <Header {...props} />
            <main>
                {children}
            </main>
        </div>
    );
}

export default UserPageContainer;