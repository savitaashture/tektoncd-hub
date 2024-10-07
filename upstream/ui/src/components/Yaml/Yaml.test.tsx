import React from 'react';
import { mount } from 'enzyme';
import Yaml from '.';

const yaml = `apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: az
  labels:
    app.kubernetes.io/version: "0.1"
  annotations:
    tekton.dev/pipelines.minVersion: "0.12.1"
    tekton.dev/tags: cli
    tekton.dev/displayName: "azure cli"
spec:
  description: >-
    This task performs operations on Microsoft Azure resources using az.

    After creating the task, you should now be able to execute az commands by
    specifying the command you would like to run as the ARGS param. The ARGS
    param takes an array of az subcommands that will be executed as part of
    this task.

  params:
  - name: az-image
    description: Azure CLI container image to run this task
    default: mcr.microsoft.com/azure-cli:2.0.78
  - name: ARGS
    type: array
    description: Azure CLI arguments to run
  steps:
  - name: az
    image: "$(params.az-image)"
    command: ["az"]
		args: ["$(params.ARGS)"]`;

it('should render the YAML file', () => {
  const component = mount(<Yaml value={yaml} />);

  expect(component.find('SyntaxHighlighter').length).toBe(1);

  expect(component.debug()).toMatchSnapshot();
});
