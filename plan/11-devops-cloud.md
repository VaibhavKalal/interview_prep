# 11 — Containers, Kubernetes, CI/CD & Cloud (Week 13, first half)

**Why it matters:** you love DevOps; this is where you can stand out. Senior backend engineers
are expected to own their services in production, not hand them over.

## Checklist

### L1 — Use
- [ ] Dockerfile basics; images vs containers; volumes; networks; docker compose
- [ ] Kubernetes objects: Pod, Deployment, Service, Ingress, ConfigMap, Secret, Job/CronJob
- [ ] CI/CD pipeline stages: build, test, scan, deploy (GitHub Actions)
- [ ] Core AWS: VPC, subnets, security groups, EC2, ECS/EKS, Lambda, RDS, S3, IAM, SQS/SNS, CloudWatch

### L2 — Internals
- [ ] What a container really is: Linux namespaces (pid, net, mnt, ...) + cgroups + a layered filesystem (overlayfs)
- [ ] Image layers and caching; multi-stage builds; distroless/scratch images (great fit for Go)
- [ ] Kubernetes control plane: API server, etcd, scheduler, controller manager; kubelet; the reconciliation loop
- [ ] Kubernetes networking: Services (ClusterIP, NodePort, LoadBalancer), kube-proxy, DNS
- [ ] Requests vs limits; CPU throttling; OOMKilled
- [ ] IAM: roles, policies, least privilege; IRSA / pod identity

### L3 — Production
- [ ] Liveness vs readiness vs startup probes, and how wrong probes cause outages
- [ ] Rolling updates, blue-green, canary; rollbacks; `PodDisruptionBudget`
- [ ] Horizontal Pod Autoscaler; scaling on custom metrics (e.g. SQS queue depth with KEDA)
- [ ] Graceful shutdown in Kubernetes: `SIGTERM`, `preStop`, `terminationGracePeriodSeconds`
- [ ] Infrastructure as Code with Terraform: state, plan/apply, modules, remote state locking
- [ ] Secrets management (AWS Secrets Manager, Vault, External Secrets)
- [ ] Supply chain basics: image scanning, pinned versions, minimal images

## Lab (weekend, in `labs/11-devops/`)

1. Containerize the Go service from module 06 with a multi-stage build into a `scratch` or distroless image. Compare image sizes with a naive build.
2. Run a local cluster with `kind` or `minikube`. Deploy it with a Deployment, Service, and readiness/liveness probes.
3. **Break it:** set a memory limit too low and watch `OOMKilled`; set a wrong readiness probe and watch traffic fail during a rollout.
4. Do a rolling update and confirm zero dropped requests (load test during the rollout) thanks to graceful shutdown.
5. Write a GitHub Actions workflow: `go vet`, `go test -race`, build, and push the image.
6. Optional: Terraform an SQS queue + DLQ in LocalStack.

## Reading
- *Kubernetes Up & Running* — Burns, Beda, Hightower: ch. 1–10
- Docker docs: "Multi-stage builds", "Best practices for writing Dockerfiles"
- Julia Evans' zines/blog on containers and namespaces
- Terraform tutorials (developer.hashicorp.com/terraform/tutorials)

## Self-test
1. What is a container, in terms of Linux kernel features?
2. Liveness vs readiness: what happens if your liveness probe checks the database?
3. Your pod restarts every few hours with exit code 137. Investigate.
4. Walk through what happens from `kubectl apply` to a running pod.
5. How do you deploy a database schema change and a code change safely together?

## Beyond the interview

- Build a container by hand with `unshare` and `chroot` (Liz Rice's "Containers from Scratch" talk walks through this in Go).
- Do "Kubernetes the Hard Way" (Kelsey Hightower) once to understand every control plane component.
- Read the AWS Well-Architected Framework's Reliability and Security pillars.
