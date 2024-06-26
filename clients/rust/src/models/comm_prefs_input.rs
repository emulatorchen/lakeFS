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
pub struct CommPrefsInput {
    /// the provided email
    #[serde(rename = "email", skip_serializing_if = "Option::is_none")]
    pub email: Option<String>,
    /// user preference to receive feature updates
    #[serde(rename = "featureUpdates")]
    pub feature_updates: bool,
    /// user preference to receive security updates
    #[serde(rename = "securityUpdates")]
    pub security_updates: bool,
}

impl CommPrefsInput {
    pub fn new(feature_updates: bool, security_updates: bool) -> CommPrefsInput {
        CommPrefsInput {
            email: None,
            feature_updates,
            security_updates,
        }
    }
}

