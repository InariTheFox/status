import React from "react";
import classNames from "classnames";
import { ReactComponent as Check } from "../assets/check.svg";
import { ReactComponent as Question } from "../assets/question.svg";

export interface IOverallStatusProps {
    state: string;
}

const OverallStatus: React.FC<IOverallStatusProps> = (props) => {

    const getStateColour = () => {
        switch (props.state) {
            case "operational":
                return "text-dark-green bg-green-400 dark:bg-green-400";
            case "degraded":
                return "text-dark-yellow bg-yellow-500 dark:bg-yellow-500";
            default:
                return "text-black bg-gray-300 dark:bg-white";
        }
    }

    const getStateIcon = () => {
        switch (props.state) {
            case "operational":
                return <Check className="w-12 h-12 -ml-1" />
            default:
                return <Question className="w-12 h-12" />
        }
    }

    const getStateText = () => {
        switch (props.state) {
            case "operational":
                return "All services are online";
            case "degraded":
                return "One or more services are degraded";
            default:
                return "Unknown"
        }
    }

    return (
        <div className="container mx-auto">
            <div className="p-5 mt-8 md:mt-24 dark:text-white rounded font-semibold flex flex-col space-y-4 items-center">
                <div className={classNames("dark:bg-white dark:text-black rounded-full flex justify-center items-center p-3 w-14 h-14", getStateColour())}>
                    {getStateIcon()}
                </div>
                <div className="text-3xl">{getStateText()}</div>
            </div>
        </div>
    );
}

export default OverallStatus;