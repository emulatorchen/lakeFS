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
pub struct ImportCreationResponse {
    /// The id of the import process
    #[serde(rename = "id")]
    pub id: String,
}

impl ImportCreationResponse {
    pub fn new(id: String) -> ImportCreationResponse {
        ImportCreationResponse {
            id,
        }
    }
}

