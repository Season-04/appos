/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  query getInstalledApplications {\n    installedApplications {\n      id\n      name\n    }\n  }\n": types.GetInstalledApplicationsDocument,
    "\n  fragment EditUserFragment on User {\n    id\n    name\n    emailAddress\n    role\n  }\n": types.EditUserFragmentFragmentDoc,
    "\n  mutation UpdateUserMutation($id: ID!, $name: String!, $role: UserRole!) {\n    updateUser(input: { id: $id, name: $name, role: $role }) {\n      ...EditUserFragment\n    }\n  }\n": types.UpdateUserMutationDocument,
    "\n  query GetUsers {\n    users {\n      id\n      name\n      emailAddress\n      role\n      lastSeenAt\n      ...EditUserFragment\n    }\n  }\n": types.GetUsersDocument,
    "\n  mutation CreateUserMutation(\n    $name: String!\n    $emailAddress: String!\n    $role: UserRole!\n    $password: String!\n  ) {\n    createUser(\n      input: {\n        name: $name\n        emailAddress: $emailAddress\n        role: $role\n        password: $password\n      }\n    ) {\n      id\n    }\n  }\n": types.CreateUserMutationDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query getInstalledApplications {\n    installedApplications {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  query getInstalledApplications {\n    installedApplications {\n      id\n      name\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment EditUserFragment on User {\n    id\n    name\n    emailAddress\n    role\n  }\n"): (typeof documents)["\n  fragment EditUserFragment on User {\n    id\n    name\n    emailAddress\n    role\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation UpdateUserMutation($id: ID!, $name: String!, $role: UserRole!) {\n    updateUser(input: { id: $id, name: $name, role: $role }) {\n      ...EditUserFragment\n    }\n  }\n"): (typeof documents)["\n  mutation UpdateUserMutation($id: ID!, $name: String!, $role: UserRole!) {\n    updateUser(input: { id: $id, name: $name, role: $role }) {\n      ...EditUserFragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query GetUsers {\n    users {\n      id\n      name\n      emailAddress\n      role\n      lastSeenAt\n      ...EditUserFragment\n    }\n  }\n"): (typeof documents)["\n  query GetUsers {\n    users {\n      id\n      name\n      emailAddress\n      role\n      lastSeenAt\n      ...EditUserFragment\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreateUserMutation(\n    $name: String!\n    $emailAddress: String!\n    $role: UserRole!\n    $password: String!\n  ) {\n    createUser(\n      input: {\n        name: $name\n        emailAddress: $emailAddress\n        role: $role\n        password: $password\n      }\n    ) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation CreateUserMutation(\n    $name: String!\n    $emailAddress: String!\n    $role: UserRole!\n    $password: String!\n  ) {\n    createUser(\n      input: {\n        name: $name\n        emailAddress: $emailAddress\n        role: $role\n        password: $password\n      }\n    ) {\n      id\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;