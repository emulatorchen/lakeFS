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
pub struct Group {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    /// Unix Epoch in seconds
    #[serde(rename = "creation_date")]
    pub creation_date: i64,
}

impl Group {
    pub fn new(id: String, creation_date: i64) -> Group {
        Group {
            id,
            name: None,
            creation_date,
        }
    }
}

