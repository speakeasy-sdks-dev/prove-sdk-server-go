# V3CompleteResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ChangeDetected`                                                   | *bool*                                                             | :heavy_check_mark:                                                 | ChangeDetected returns true if the individual information changed. | false                                                              |
| `Next`                                                             | map[string]*string*                                                | :heavy_check_mark:                                                 | Next contains the next set of allowed calls in the same flow.      | {<br/>"done": null<br/>}                                           |
| `Success`                                                          | *bool*                                                             | :heavy_check_mark:                                                 | Success returns true if the individual was verified successfully.  | true                                                               |