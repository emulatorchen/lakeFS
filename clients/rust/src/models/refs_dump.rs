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
pub struct RefsDump {
    #[serde(rename = "commits_meta_range_id")]
    pub commits_meta_range_id: String,
    #[serde(rename = "tags_meta_range_id")]
    pub tags_meta_range_id: String,
    #[serde(rename = "branches_meta_range_id")]
    pub branches_meta_range_id: String,
}

impl RefsDump {
    pub fn new(commits_meta_range_id: String, tags_meta_range_id: String, branches_meta_range_id: String) -> RefsDump {
        RefsDump {
            commits_meta_range_id,
            tags_meta_range_id,
            branches_meta_range_id,
        }
    }
}

