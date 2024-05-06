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
pub struct GarbageCollectionConfig {
    /// Duration in seconds. Objects created in the recent grace_period will not be collected.
    #[serde(rename = "grace_period", skip_serializing_if = "Option::is_none")]
    pub grace_period: Option<i32>,
}

impl GarbageCollectionConfig {
    pub fn new() -> GarbageCollectionConfig {
        GarbageCollectionConfig {
            grace_period: None,
        }
    }
}
