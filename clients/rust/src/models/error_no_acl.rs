/*
 * lakeFS API
 *
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: services@treeverse.io
 * Generated by: https://openapi-generator.tech
 */

use crate::models;

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ErrorNoAcl {
    /// short message explaining the error
    #[serde(rename = "message")]
    pub message: String,
    /// true if the group exists but has no ACL
    #[serde(rename = "no_acl", skip_serializing_if = "Option::is_none")]
    pub no_acl: Option<bool>,
}

impl ErrorNoAcl {
    pub fn new(message: String) -> ErrorNoAcl {
        ErrorNoAcl {
            message,
            no_acl: None,
        }
    }
}

