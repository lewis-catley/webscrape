import React from "react";
import { Switch, Route } from "react-router-dom";
import { Home, Result } from "../pages";

export const Routes: React.FC = () => {
  return (
    <Switch>
      <Route exact path="/" component={Home} />
      <Route path="/results/:id" component={Result} />
    </Switch>
  );
};
