Test the api:

```bash
cat api-test/get_events.txt | vegeta attack -duration=10s -rate=50 | tee results.bin | vegeta report
```
